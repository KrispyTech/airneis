import { Text, TouchableHighlight, View } from "react-native"
import { locales } from "../utils/constants/translations"
import i18n from "i18next"
import { useTranslation } from "react-i18next"
import { clsx } from "clsx"
import { ChevronDownIcon, ChevronRightIcon } from "react-native-heroicons/solid"
import mobileConfig from "../../mobileConfig"
import { useState } from "react"

const LanguageSelector = props => {
  const { className } = props
  const { t } = useTranslation()
  const [isDeployed, setIsDeployed] = useState(false)
  const handleLanguageChange = locale => () => {
    i18n.changeLanguage(locale)
  }

  return (
    <View className={clsx(className, "border-2 border-black w-64 mx-auto px-1 rounded-lg duration-500")}>
      <TouchableHighlight onPress={() => setIsDeployed(!isDeployed)}>
        <View className="flex flex-row items-center justify-between px-3">
          <ChevronDownIcon {...mobileConfig.icons.mdIcon} color="black" />

          <Text className="text-lg mb-3 mt-3">{t("translations.selectLanguage")}</Text>
        </View>
      </TouchableHighlight>
      {isDeployed && (
        <View className="flex gap-3 mb-4">
          {locales.map(locale => (
            <TouchableHighlight onPress={handleLanguageChange(locale)} key={locale}>
              <View className="flex flex-row items-center border border-black mx-4 px-2 py rounded">
                <ChevronRightIcon {...mobileConfig.icons.mdIcon} color="black" />
                <Text className="text-xl ml-2">{t(`translations.${locale}`)}</Text>
              </View>
            </TouchableHighlight>
          ))}
        </View>
      )}
    </View>
  )
}

export default LanguageSelector
