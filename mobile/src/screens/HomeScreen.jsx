import { Button, ScrollView, Text } from "react-native"
import Page from "../components/Page"
import screens from "../utils/constants/screens"
import { getCategories } from "../utils/services/categoriesServices"
import { useEffect, useState } from "react"
import ErrorScreen from "./ErrorScreen"
import errors from "../utils/constants/errors"

const HomeScreen = props => {
  const { navigation } = props
  const [categories, setCategories] = useState([])
  const [hasError, setError] = useState(false)
  const navigateToCategory = id => () => navigation.navigate(screens.categoryPage, { id })
  const navigateToSignUp = () => navigation.navigate(screens.signUpScreen)

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
    return <ErrorScreen navigation={navigation} text={errors.homePage.unableToLoadCategories} />
  }

  return (
    <Page navigation={navigation}>
      <ScrollView>
        <Text>Home page</Text>
        <Button title="SignUp" onPress={navigateToSignUp}>
          Sign up
        </Button>
        {categories.map(({ name, ID: id }) => (
          <Button title={name} onPress={navigateToCategory(id)} key={id} />
        ))}
      </ScrollView>
    </Page>
  )
}

export default HomeScreen
