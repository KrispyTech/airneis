import { Text } from "react-native"
import Page from "../components/Page"

const CategoriesScreen = props => {
  const { navigation } = props

  return (
    <Page navigation={navigation}>
      <Text>Category Page</Text>
    </Page>
  )
}

export default CategoriesScreen
