#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
Generate a ransomware-safe B2 API key that does not have the power to destroy files.
"""
import os,subprocess

KEYNAME = 'safe'
BUCKETNAME = '<bucketName>'

env = os.environ.copy()
env['B2_APPLICATION_KEY_ID'] = '<keyID>'
env['B2_APPLICATION_KEY'] = '<secretKey>'


capabilities = [
    "listAllBucketNames",
    "listBuckets",
    "listFiles",
    "readBucketEncryption",
    "readBucketReplications",
    "readBucketRetentions",
    "readBuckets",
    "readFileLegalHolds",
    "readFileRetentions",
    "readFiles",
    "shareFiles",
    "writeBuckets",
    "writeFiles"
]

cmd = ['./b2','key','create','--bucket',BUCKETNAME,KEYNAME,','.join(capabilities)]
subprocess.call(cmd,env=env)
