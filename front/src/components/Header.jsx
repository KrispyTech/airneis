import { MagnifyingGlassIcon, ShoppingCartIcon, Bars3Icon } from "@heroicons/react/24/outline"
import webConfig from "@@/webConfig"
import Link from "next/link"
import { backoffice } from "@/constants"
import routes from "@/routes"
import { useTranslation } from "next-i18next"
import { useRouter } from "next/router"

const Header = ({ page }) => {
  const { t } = useTranslation("common")
  const { locales, push } = useRouter()
  const changeLocale = locale => {
    push("", "", { locale })
  }

  return (
    <header className="bg-secondary text-primary px-4 flex justify-between items-center sticky top-0 z-30">
      <Link href={routes.home} className="text-xl md:text-4xl py-3 font-extrabold">
        ÀIRNEIS
      </Link>
      {page !== backoffice ? (
        <nav>
          <ul className="flex flex-row items-center gap-5">
            <li>
              <select className="bg-secondary " onInput={({ target: { value } }) => changeLocale(value)}>
                <option value="">{t("header.selectLanguage")}</option>
                {locales.map(locale => (
                  <option key={locale} value={locale}>
                    {t(`header.locales.${locale}`)}
                  </option>
                ))}
              </select>
            </li>
            <li>
              <MagnifyingGlassIcon {...webConfig.icons.lgIcon} />
            </li>
            <li>
              <ShoppingCartIcon {...webConfig.icons.lgIcon} />
            </li>
            <li>
              {/* As this icon is smaller that the other we need it to have a bigger size */}
              <Bars3Icon {...webConfig.icons.xlIcon} />
            </li>
          </ul>
        </nav>
      ) : (
        <p className="text-xl md:text-4xl font-extrabold">BACKOFFICE</p>
      )}
    </header>
  )
}

export default Header
