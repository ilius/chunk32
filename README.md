

# Encode a file (in Linux/Unix)
Type `cat FILE_PATH | chunk32`.
Example:

    $ cat ~/roses.txt | chunk32
    a9qq 6sbk 41gq 4s90
    e9jp 8b0a asmp yv35
    ehsj 0rbj cmg6 4v3n
    cmp0 mwvn cxgq 4839
    ecg7 6xv5 cnt2 r2j1
    dsj2 0wvf 41gq 4s90
    1wpy x9e-

# Encode text from Standard input (in Linux/Unix)
Type `chunk32` and Enter, then type/paste your text, then press Control+D

# Encrypt and encode text from Standard input (in Linux/Unix)
Type `gpg -c | chunk32`, then it will prompt for a password and your need to enter your desired password (twice), then type/paste your text, then press Enter and then Control+D.
The output should be longer than without encryption.
Example:

    $ gpg -c | chunk32 
    Roses are red,     
    Violets are blue,
    sugar is sweet,
    And so are you.
    [Control+D]
    hg6g 81r3 0b0z p21s
    7enx adk0 t9qg 3dqw
    9edb xmbk d8xw 6hvy
    etww 1nax t9n4 z0zj
    zwhq h6ay w4vj 72pw
    mj0n y1qx x810 dj15
    nata wwwq tpac m4m4
    nxwt 2z1x db1b 4n0q
    8hvj 32dc 7s0m 066d
    ct4r g37n d29x hypn
    9k24 297t 7s9e z5h9
    hw48 ask3 zjva p04w
    2xgk m9nx rayd c---

# Decode from Standard input (in Linux/Unix)
Type `chunk32 -d` and Enter, then type/paste the chunk32-encoded text, then press Control+D

# Decode and decript from Standard input (in Linux/Unix)
Type `chunk32 -d | gpg -d`, then type/paste your text, then press Enter and then Control+D, then it will prompt for a password and your need to enter the password and press Enter, then the final text/binary data will be displayed. If it's binary, you may want to pipe it again `chunk32` (like `chunk32 -d | gpg -d | chunk32`) or `base64` or for example (Cryptocurrency private keys usually use base58)

# Check Character
Check character can be enabled with `-check` flag when encoding.
Check character is one extra character at the end of each line of encoded text, which is the signature of that line (chunk of data) that can help identifying possible errors (in typing or storing the encoded data)
For example:

    $ cat ~/roses.txt | chunk32 -check
    a9qq 6sbk 41gq 4s90 6
    e9jp 8b0a asmp yv35 ~
    ehsj 0rbj cmg6 4v3n g
    cmp0 mwvn cxgq 4839 0
    ecg7 6xv5 cnt2 r2j1 h
    dsj2 0wvf 41gq 4s90 g
    1wpy x9e- h

When decoding, you never need to pass `-check` flag. We will detect wheather or not the encoded text has check character.

# Use uppercase letters
If you want to encode with uppercase letters, pass `-u` flag.

When docoding, lowercase and uppercase letters are the same (so you can even mix them)
