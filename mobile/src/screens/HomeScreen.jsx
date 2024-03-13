import { Button, Text } from "react-native"
import Page from "../components/Page"
import screens from "../utils/constants/screens"

const HomeScreen = props => {
  const { navigation } = props
  const goToCategory = () => navigation.navigate(screens.categoryPage)

  return (
    <Page navigation={navigation}>
      <Text>Home page</Text>
      <Button title="Go to categories" onPress={goToCategory} />
    </Page>
  )
}

export default HomeScreen
