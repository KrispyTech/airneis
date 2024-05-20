import { useRouter } from "next/router"
import { useTranslation } from "next-i18next"

const SelectLanguage = () => {
  const { t } = useTranslation("common")
  const { locales, push, pathname, locale, query, asPath } = useRouter()
  const changeLocale = async newLocale => {
    await push({ pathname, query }, asPath, { locale: newLocale })
  }

  const handleInput = async ({ target: { value } }) => {
    await changeLocale(value)
  }

  return (
    <select className="bg-secondary " onInput={handleInput}>
      <option value={locale}>{t("header.selectLanguage")}</option>
      {locales.map(currentLocale => (
        <option key={currentLocale} value={currentLocale}>
          {t(`header.locales.${currentLocale}`)}
        </option>
      ))}
    </select>
  )
}

export default SelectLanguage
