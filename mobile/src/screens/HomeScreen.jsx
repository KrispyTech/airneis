import { Button, ScrollView, Text, View } from "react-native"
import Page from "../components/Page"
import screens from "../utils/constants/screens"
import { getCategories } from "../utils/services/categoriesServices"
import { useEffect, useState } from "react"
import ErrorScreen from "./ErrorScreen"
import errors from "../utils/constants/errors"
import { useTranslation } from "react-i18next"
import LanguageSelector from "../components/LanguageSelector"
import useSessionStore from "../stores/sessions"
import { destroySession } from "../utils/services/sessionServices"

const HomeScreen = props => {
  const { navigation } = props
  const [categories, setCategories] = useState([])
  const [hasError, setError] = useState(false)
  const { isLogged, signOut } = useSessionStore()
  const navigateToCategory = id => () => navigation.navigate(screens.categoryPage, { id })
  const { t } = useTranslation()
  const navigateToSignUp = () => navigation.navigate(screens.signUpScreen)
  const navigateToSignIn = () => navigation.navigate(screens.signInScreen)
  const handleSinOut = async () => {
    await destroySession()
    signOut()
  }

  useEffect(() => {
    ;(async () => {
      try {
        const categoriesResult = await getCategories()
        setCategories(categoriesResult.data)
      } catch (e) {
        setError(true)
      }
    })()
  }, [])

  if (hasError) {
    return <ErrorScreen navigation={navigation} text={t(errors.homePage.unableToLoadCategories)} />
  }

  return (
    <Page navigation={navigation}>
      <ScrollView>
        <Text>{t("homePage.title")}</Text>
        <LanguageSelector className="mt-5" />
        {isLogged ? (
          <Button title={t("homePage.signOut")} onPress={handleSinOut} />
        ) : (
          <View className="flex flex-row justify-center">
            <Button title={t("homePage.signUp")} onPress={navigateToSignUp} />
            <Button title={t("homePage.signIn")} onPress={navigateToSignIn} />
          </View>
        )}
        {categories.map(({ name, ID: id }) => (
          <Button title={name} onPress={navigateToCategory(id)} key={id} />
        ))}
      </ScrollView>
    </Page>
  )
}

export default HomeScreen
