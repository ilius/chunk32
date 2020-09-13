# What is it for?
The use case is that you have a binary data, possibly an encrypted text message, and you want for write it down manually on a paper (for example) because it's *so sensitive* that you don't want to store it on your computer (or you just want the paper as backup in case your computer / hard drive was ever broken or stolen), and you definitely *can't risk exposing that data to a printer device* (specially one that is owned by somebody else).

Another case is that you want to read this binary data (or encrypted text message) over a telephone line (that is not very secure).

I made a character/text encoding special for these kinds of cases.

# How does it work?
We use [Crockford's Base32](https://www.crockford.com/base32.html) encoding which is more suitable than other Base32 variants (because it avoids using both of similar characters like `I`, `L` and one, or `O` and zero). And we specially avoid Base58 or Base64 because they contain more visually similar letters, and they use both uppercase and lowercase letters, which makes it much harder to duplicate manually (read and write down) or read to someone else (in person or over phone).

Although Base32-encoded string is ~%20 longer than Base64-encoded and ~%17 longer than Base58-encoded (for large inputs), but using Crockford's Base32 saves a lot of time for the mentioned purposes (you can test and see)

Since Base32 maps each 5 bits to a 8 bits, so it maps each 5 bytes to 8 bytes.
And 16 characters is a reasonable number of characters to use in every row/line (it fits in a notebook paper for example).

So first we split the input binary data into chunks of 10 bytes, then encode each chunk with Crockford's Base32, and each encoded text becomes a line.

The benefit of splitting into chunks (each becoming a line), is that we can take advantage of **Check Symbol** supported by [Crockford's Base32](https://www.crockford.com/base32.html), by adding the **Check Symbol** of each chunk / line to the end of that line, as shown in [examples below](#check-character). So that if you mistype / misread a few characters in each row (16 characters), the check symbol will almost certainly mismatch (very low possibility it will not) and you would know. Specially if the data is compressed or encrypted, this will prevent the wrong data going through the next channel (decompression or decryption) and save you some time. And even if the receiver does not validate the check symbol until it's too late (the original data is not available) and encoded data is broken, the check symbol gives you the chance to try and find the wrong characters and fix then, assuming they are no more than a few in the broken row/chunk. But if you use a checksum (like md5 or sha1) for the entire input (instead of chunking and check symbol), fixing the broken data would be much harder (might be impossible for large input). Generally the purpose of overall md5/sha1 checksum is only validation, not fixing. You can still use them alongside `chunk32` as a second safety measure.


# How to install it
First [install Go](https://golang.org/doc/install), then open a command line and run:

    go get github.com/ilius/chunk32


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

# Decode and decrypt from Standard input (in Linux/Unix)
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

When decoding, you never need to pass `-check` flag. We will detect whether or not the encoded text has check character.

# Use uppercase letters
If you want to encode with uppercase letters, pass `-u` flag.

When decoding, lowercase and uppercase letters are the same (so you can even mix them)
