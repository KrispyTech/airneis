import { MagnifyingGlassIcon, ShoppingCartIcon, Bars3Icon } from "@heroicons/react/24/outline"
import webConfig from "@@/webConfig"
import Link from "next/link"
import { backoffice } from "@/constants"
import routes from "@/routes"
import { useTranslation } from "next-i18next"
import SelectLanguage from "@/components/ui/SelectLanguage"

const Header = ({ page }) => {
  const { t, locale } = useTranslation("common")

  return (
    <header className="bg-secondary text-primary px-4 flex justify-between items-center sticky top-0 z-30">
      <Link href={routes.home} className="text-xl md:text-4xl py-3 font-extrabold" locale={locale}>
        ÀIRNEIS
      </Link>
      {page !== backoffice ? (
        <nav>
          <ul className="flex flex-row items-center gap-5">
            <SelectLanguage />
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
        <div>
          <SelectLanguage />
          <p className="text-xl md:text-4xl font-extrabold">{t("header.backOffice")}</p>
        </div>
      )}
    </header>
  )
}

export default Header
