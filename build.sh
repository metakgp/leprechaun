git branch -D prod;
git checkout -b prod;
sed -ie "s/dist//g" .gitignore;
gulp;
git add .;
git commit -m "build latest master";
git push heroku -f prod:master;
git checkout master;
