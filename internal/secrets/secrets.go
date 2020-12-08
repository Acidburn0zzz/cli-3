package secrets

import (
	"errors"

	"github.com/ActiveState/cli/internal/keypairs"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
	secretsapi "github.com/ActiveState/cli/pkg/platform/api/secrets"
	secretsapiClient "github.com/ActiveState/cli/pkg/platform/api/secrets/secrets_client/secrets"
	secretsModels "github.com/ActiveState/cli/pkg/platform/api/secrets/secrets_models"
	"github.com/ActiveState/cli/pkg/platform/authentication"
	"github.com/ActiveState/cli/pkg/platform/model"
)

// Save will add a new secret for this user or update an existing one.
func Save(secretsClient *secretsapi.Client, encrypter keypairs.Encrypter, org *mono_models.Organization, project *mono_models.Project,
	isUser bool, secretName, secretValue string) error {

	logging.Debug("attempting to upsert user-secret for org=%s", org.OrganizationID.String())
	encrStr, err := encrypter.EncryptAndEncode([]byte(secretValue))
	if err != nil {
		return err
	}

	params := secretsapiClient.NewSaveAllUserSecretsParams()
	params.OrganizationID = org.OrganizationID
	secretChange := &secretsModels.UserSecretChange{
		Name:   &secretName,
		Value:  &encrStr,
		IsUser: &isUser,
	}
	if project != nil {
		secretChange.ProjectID = project.ProjectID
	}

	params.UserSecrets = append(params.UserSecrets, secretChange)

	_, err = secretsClient.Secrets.Secrets.SaveAllUserSecrets(params, authentication.Get().ClientAuth())
	if err != nil {
		logging.Error("error saving user secret: %v", err)
		return locale.WrapError(err, "secrets_err_save", "", err.Error())
	}

	return nil
}

// ShareWithOrgUsers will share the provided secret with all other users in the organization
// who have a valid public-key available.
func ShareWithOrgUsers(secretsClient *secretsapi.Client, org *mono_models.Organization, project *mono_models.Project, secretName, secretValue string) error {
	currentUserID, err := secretsClient.AuthenticatedUserID()
	if err != nil {
		return err
	}

	members, err := model.FetchOrgMembers(org.URLname)
	if err != nil {
		return err
	}

	for _, member := range members {
		if currentUserID != member.User.UserID {
			pubKey, err := keypairs.FetchPublicKey(secretsClient, member.User)
			if err != nil {
				if errors.Is(err, keypairs.ErrKeypairNotFound) {
					logging.Info("User `%s` has no public-key", member.User.Username)
					// this is okay, just do what we can
					continue
				}
				return err
			}

			ciphertext, err := pubKey.EncryptAndEncode([]byte(secretValue))
			if err != nil {
				logging.Error("Encryptying secret `%s` for user `%s`: %s", secretName, member.User.Username)
				// this is a local issue with the user's keys, so we try and move on
				continue
			}

			share := &secretsModels.UserSecretShare{
				Name:  &secretName,
				Value: &ciphertext,
			}
			if project != nil {
				share.ProjectID = project.ProjectID
			}

			err = secretsapi.SaveSecretShares(secretsClient, org, member.User, []*secretsModels.UserSecretShare{share})
			if err != nil {
				// a potentially unrecoverable err, so we stop here
				return err
			}
			logging.Info("Update secret `%s` for user `%s`", secretName, member.User.Username)
		}
	}

	return nil
}

func LoadKeypairFromConfigDir() (keypairs.Keypair, error) {
	kp, err := keypairs.LoadWithDefaults()
	if err != nil {
		return nil, err
	}
	return kp, nil
}

// DefsByProject fetches the secret definitions for the current user relevant to the given project
func DefsByProject(secretsClient *secretsapi.Client, owner string, projectName string) ([]*secretsModels.SecretDefinition, error) {
	pjm, fail := model.FetchProjectByName(owner, projectName)
	if fail != nil {
		return nil, fail
	}

	return secretsapi.FetchDefinitions(secretsClient, pjm.ProjectID)
}

// ByProject fetches the secrets for the current user relevant to the given project
func ByProject(secretsClient *secretsapi.Client, owner string, projectName string) ([]*secretsModels.UserSecret, error) {
	result := []*secretsModels.UserSecret{}

	pjm, fail := model.FetchProjectByName(owner, projectName)
	if fail != nil {
		return result, fail
	}

	org, fail := model.FetchOrgByURLName(owner)
	if fail != nil {
		return result, fail
	}

	secrets, fail := secretsapi.FetchAll(secretsClient, org)
	if fail != nil {
		return result, fail
	}

	for _, secret := range secrets {
		if secret.ProjectID != pjm.ProjectID {
			continue // We only want secrets belonging to our project
		}
		result = append(result, secret)
	}

	return result, nil
}
