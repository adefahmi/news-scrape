# news-scrape
BMKG Siaran pers scrape

# Install
It's just a single binary file, no external dependencies. Just download the appropriate version of [executable from latest release for your OS](https://github.com/adefahmi/news-scrape/releases). Then rename and give it execute permission.
```
mv scrape news-scrape  
sudo chmod +x news-scrape
```
If you want to install it globally (run from any directory of your system), put it in your systems $PATH directory.
```
sudo mv news-scrape /usr/local/bin/news-scrape
```
Done!

# How to use
```
news-scrape
```
app running on port 4000
you can access api from path [http://localhost:4000/news](http://localhost:4000/news)
