import os
import sys
import re
import subprocess
import shutil
import uuid
import tempfile

import helpers


class TestUpdates(helpers.IntegrationTest):

    def __init__(self, *args, **kwargs):
        super(TestUpdates, self).__init__(*args, **kwargs)
        self.constants = helpers.get_constants()

    def setUp(self):
        self.clear_config() # clear_config will set up env vars, so below we ensure the proper ones are set
        self.env["ACTIVESTATE_CLI_AUTO_UPDATE_TIMEOUT"] = "10" # Ensure auto-update has plenty of time to run
        self.env["ACTIVESTATE_CLI_UPDATE_BRANCH"] = "master" # Our own branch probably doesn't have update bits yet

        self.temp_path = self.get_temp_path()
        os.mkdir(self.temp_path)
        shutil.copy(os.path.join(self.test_dir, "testdata", "activestate.yaml"), self.temp_path)
        self.set_cwd(self.temp_path)

    def get_version(self, temp_bin):
        output = self.spawn_command_blocking("%s --version" % temp_bin)
        version_regex = re.compile(".*(\d\.\d\.\d-\d{4})")
        match = version_regex.search(str(output))
        if match:
            return match.group(1)

    def assert_version_match(self, same, temp_bin, message):
        assertFn = self.assertEqual if same else self.assertNotEqual
        assertFn(self.get_version(temp_bin), self.constants["Version"], message)

    def test_auto_update_works(self):
        temp_bin = self.get_temp_bin()
        self.env["ACTIVESTATE_CLI_DISABLE_UPDATES"] = "false"
        self.assert_version_match(False, temp_bin, "Version number should not match because auto-update should have occurred")
        self.wait()

    def test_manual_update_works(self):
        temp_bin = self.get_temp_bin()
        self.spawn_command("%s update" % temp_bin)
        self.expect("Update completed")
        self.wait()

        self.assert_version_match(False, temp_bin, "Version number should not match because we ran update")

    def test_locked_version_works(self):
        self.env["ACTIVESTATE_CLI_DISABLE_UPDATES"] = "false"
        temp_bin = self.get_temp_bin() 
        self.spawn_command("%s update --lock" % temp_bin)
        self.expect("Version locked at")
        self.wait()

if __name__ == '__main__':
    helpers.Run()
