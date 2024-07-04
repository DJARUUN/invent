# Invent

#### **inv**ert + ind**ent** = **invent**

A very serious formatter that does what everyone has been missing. OUTdenting your code!
Gone are the days of code inside of code being indented two or four or whatever spaces. 

This tool lets you do the exact opposite. It takes any file or directory you specify and inverts it's indentation. Or should I say, outdents your code!

For me this has been the most beneficial change made in my projects boosting both productivity and throughput.

### Installation / usage

Just download the binary for your operating system in the releases.
If it says `Permission denied` you have to run `chmod +x invent` to give it executable permissions.

### Tip

Do not and I mean NOT run this on a directory with a bunch of files that aren't plain text files. There are some attempts to make it ignore those types of files but it may still find it and it will not be a good day for anyone when it tries to outdent an executable file, for example.

From my small testing this *should* not happen but it very much does sometimes so just be careful to not actually mess anything up.

Also don't run this on any files of intendation based languages because it will ofcourse break them completely, and just running it again doesn't really fix it to what it was (I feel like it should but it doesn't).