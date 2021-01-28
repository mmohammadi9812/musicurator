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
go build ./cmd/musicurator/
```
### Structure
folders are structured like this:  
+ **./core/**: core files to app, which includes search and renaming functionality
+ **./cmd/**: package for cmd tool
+ **./gui/**: