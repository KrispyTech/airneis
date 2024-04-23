import { Text, TouchableHighlight, View } from "react-native"
import { locales } from "../utils/constants/translations"
import i18n from "i18next"
import { useTranslation } from "react-i18next"
import { clsx } from "clsx"
import { ChevronRightIcon } from "react-native-heroicons/solid"
import mobileConfig from "../../mobileConfig"

const LanguageSelector = props => {
  const { className } = props
  const { t } = useTranslation()
  const handleLanguageChange = locale => () => {
    i18n.changeLanguage(locale)
  }

  return (
    <View className={clsx(className, "border-2 border-black w-64 mx-auto py-2 px-1 rounded-lg")}>
      <Text className="text-lg mb-3">{t("translations.selectLanguage")}</Text>
      <View className="flex gap-3">
        {locales.map(locale => (
          <TouchableHighlight onPress={handleLanguageChange(locale)} key={locale}>
            <View className="flex flex-row items-center border border-black">
              <ChevronRightIcon {...mobileConfig.icons.mdIcon} color="black" />
              <Text className="text-xl ml-2">{t(`translations.${locale}`)}</Text>
            </View>
          </TouchableHighlight>
        ))}
      </View>
    </View>
  )
}

export default LanguageSelector
