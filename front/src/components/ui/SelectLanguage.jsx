import { useRouter } from "next/router"
import { useTranslation } from "next-i18next"

const SelectLanguage = () => {
  const { t } = useTranslation("common")
  const { locales, push } = useRouter()
  const changeLocale = locale => {
    push("", "", { locale })
  }

  const handleInput = ({ target: { value } }) => {
    changeLocale(value)
  }

  return (
    <select className="bg-secondary " onInput={handleInput}>
      <option value="">{t("header.selectLanguage")}</option>
      {locales.map(locale => (
        <option key={locale} value={locale}>
          {t(`header.locales.${locale}`)}
        </option>
      ))}
    </select>
  )
}

export default SelectLanguage
