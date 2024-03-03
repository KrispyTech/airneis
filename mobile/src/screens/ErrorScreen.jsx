import Page from "../components/Page"
import { Text } from "react-native"

const ErrorScreen = props => {
  const { text, navigation } = props

  return (
    <Page navigation={navigation}>
      <Text className="text-center text-red-500 text-lg font-semibold py-5">An error occured: {text}</Text>
    </Page>
  )
}

export default ErrorScreen
