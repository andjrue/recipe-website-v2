# recipe-website-v2

### Project Setup:

**Languages:**

I'm going to use Go for the backend, and React for the front end. I really have no idea how React works, but it should be fun to learn. 

I'm starting to understand Go and see this as a great way to boost my skills. 

**Database:**

For now I'm thinking about using Mongo. I initially tried this project using Postgres and found that the DB alterations were tedious and time consuming. I also want to store Users and their recipes in a particular way, and think that Mongo would be perfect for it. If all the information could be in a single User struct, life would be good. 

**Front-End Overview:**

The front end will be simple. This won't be an application that has thousands of concurrent users (hopefully). Because of that, we can really focus on what it needs to have as opposed to what users want. 

It will consist of a login screen that querys the DB to ensure the email and password entered matches the email and hashed password we have stored. After that, we'll use their auth token for any and all operations. It would also be a good idea to set some kind of role for users. I would give myself admin priviledges in case someone comes into the website and runs amuck. Everybody else can have basic user perms. 

These perms will include being able to update/delete recipes created by that user. We don't want someone to accidentally delete someone elses recipes. 

*Layout:*

As mentioned above, the layout will be simple. After someone has logged in, they will see every recipe uploaded to the website, who uploaded it, etc. These will be clickable cards if the user sees something interesting. 

On the left side we will have a nav bar with a few options. Right now, I think we should have a users page, an account page, and a link back to the main recipe landing page. 

<ins>User Page:</ins>

This page will show every single user that's on the website. It will have their username, but not their email. We don't want someone going rogue and farming information. The email is set specifically for password reset purposes. 

These will also be in cards (I'm sensing a theme). If someone clicks into a users card, they'll be able to see every recipe they've uploaded. If something catches their eye, those cards will also be clickable. 

<ins>Account Page</ins>

This will contain user account information. If they want to update their password here, they can do that. if they want to view their recipes, they can do that. 

I think in the account page, we should open a different nav bar that separates these two things. If someone has a ton of recipes uploaded, it could get confusing. 

<ins>Recipe Landing Page</ins>

I want this to be available regardless of where you are on the website. The main page should always be available. 


### Project Workflow:

**Step 1:**

We're starting with user login and authentication. 

Each user will need a unique email and their own password. I still need to decide on specific password requirements. Probably at least 8 characters and a special 
character. 

Users will also be given a unique token that will be hashed several times. The token will expire after 12 hours, prompting them to log back in. 

Once we have users able to sign in and validated through their token, we can move on to step 2.

*Sample User Struct:*

type User Struct {
    Email string
    Password string (will be hashed)
    Recipes []Recipe (array of recipes that will be stored with the user.
}

**Step 2:**

Setting up the DB connection. 

I've never used Mongo before, so this should be fun. Initially, we'll need to write Users to the db. Once that's done, we can move on to adding recipes. 

**Step 3:**

Recipe config. 

Once we have user creation finished, we can move on to the Recipe portion of this.

This should be pretty easy to do, and will be input by users through a form. This will require its own endpoint. Probably "{endpoint}/recipes". 

Recipes will also be their own struct, which will be included in Users as shown above. 

*Sample Recipe Struct:*

type Recipe struct {
    Title string
    Description string
    Ingredients string
    TimeToMake string
    Steps string
}

That isn't exactly what it will look like, but it should go along those lines. 

We can use the auth token assigned to a user to look up who is logged in (we could also use email, but that feels slightly insecure) and assign those recipes to whoever added them. 

This way, if someone wants to see every recipe they've uploaded its much easier to handle on the front end. It's as simple as looking up their email in the DB and looping through the recipes. 

**Step 4:**

I think this is where we'll start work on the front end. This will still be far from complete, and there will be more work to do on the backend, but this feels like a good entry point for React. We'll have user and recipe creation working, it would be a good point in time to see what we can create. 

There is still more to be done after this, but I'm not exactly sure how we're going to do it. I'll add more to the readme as we get further into this. 


