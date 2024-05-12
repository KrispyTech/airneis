import { useRouter } from "next/router"
import { useTranslation } from "next-i18next"

const SelectLanguage = () => {
  const { t } = useTranslation("common")
  const { locales, push, pathname, locale, query, asPath } = useRouter()
  const changeLocale = async locale => {
    await push({ pathname, query }, asPath, { locale: locale })
  }

  const handleInput = async ({ target: { value } }) => {
    await changeLocale(value)
  }

  return (
    <select className="bg-secondary " onInput={handleInput}>
      <option value={locale}>{t("header.selectLanguage")}</option>
      {locales.map(locale => (
        <option key={locale} value={locale}>
          {t(`header.locales.${locale}`)}
        </option>
      ))}
    </select>
  )
}

export default SelectLanguage
