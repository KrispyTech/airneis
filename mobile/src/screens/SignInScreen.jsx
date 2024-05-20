import { Text } from "react-native"
import Page from "../components/Page"

const SignInScreen = props => {
  const { navigation } = props

  return (
    <Page navigation={navigation}>
      <Text>Sign In Screen</Text>
    </Page>
  )
}

export default SignInScreen
