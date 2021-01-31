Musicurator
===========
Musicurator searches your computer for music files and tries to use id3 tags to change filenames
### why
better searches through files. If you've worked with downloaded mp3 files, you may have seen odd filenames like:  
`000_will_tay.mp3`  
It's hard to search for willow song from taylor swift with names like this.  
So with `musicurator` you can fix this to: `taylor_swift__willow.mp3`  
### Installing
executable files for linux and windows can be found in `release` tab of mirrored github page  

alternatively, you may build command-line tool with:  
```bash
git clone github.com/mmohammadi9812/musicurator
cd musicurator
go build ./gui/musicurator/ # or ./cmd/musicurator/ if you want the cmd binaries
```
### Structure
folders are structured like this:  
+ **./core/**: core files to app, which includes search and renaming functionality
+ **./cmd/**: package for cmd tool
+ **./gui/**: package for gui app

please *NOTE* that cmd binaries and gui app do the same thing and are totally independent
You may use which 

### LICENSE
This repo is licensed under BSD 3-Clause License
you can find it in `LICENSE` file
