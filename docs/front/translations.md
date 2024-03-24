# Translations

To make the site available in different languages, we use
the [next-i18n](https://github.com/i18next/next-i18next/) packag

# How to use it

You can follow [this tutorial](https://www.youtube.com/watch?v=H9O9HdKNytc&ab_channel=CulesCoding) to know a bit more
this package.

### Translations files scaffolding

The translations are stored in JSON files. Those files are located in the `public/locales` folder.
You need to have one subdirectory per language and one file named `common.json` (that's the best practice to name it
like that) per directory.
Create a key per page / component and add a key per sentence you want to translate. This key should be the sentence in
camelCase.

Check the example below:

```json
{
  "home": {
    "title": "Welcome to my website",
    "description": "This is a description of my website"
  }
}
```

### How to use it in a component

First you need to import the translation from the server to your component.
You can use the following code to do that:

```javascript
import {serverSideTranslations} from "next-i18next/serverSideTranslations"


export const getServerSideProps = async ({locale}) => ({
    props: {
        ...(await serverSideTranslations(locale, ["common"]))
    }
})
```

Then you need you can use the following code in your component to use the translation:

```javascript
 const Home = () => {
    const {locale} = useRouter()
    const {t} = useTranslation("common")


    return (
        <Page className="h-screen text-secondary">
            <h2 className="text-2xl font-semibold py-5 px-3"> {t("homePage.wip")}</h2>
            <Link href="/" locale={locale}>{t("homePage.linkExample")}</Link>
        </Page>
    )
}
```

### IMPORTANT

**As the locale is part of the url you need to specify it in the Link like the example above.**