import Page from "../components/Page"
import { Text } from "react-native"
import { useTranslation } from "react-i18next"

const ErrorScreen = props => {
  const { t } = useTranslation()
  const { text, navigation } = props

  return (
    <Page navigation={navigation}>
      <Text className="text-center text-red-500 text-lg font-semibold py-5">
        {t("error.message")} {text}
      </Text>
    </Page>
  )
}

export default ErrorScreen
