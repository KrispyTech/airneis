import { MagnifyingGlassIcon, ShoppingCartIcon, Bars3Icon } from "@heroicons/react/24/outline"
import webConfig from "@@/webConfig"
import Link from "next/link"
import { backoffice } from "@/constants"

const Header = ({ page }) => (
  <header className="bg-secondary text-primary px-4 flex justify-between items-center sticky top-0 z-30">
    <Link href="/" className="text-xl md:text-4xl py-3 font-extrabold">
      ÀIRNEIS
    </Link>
    {page !== backoffice ? (
      <nav>
        <ul className="flex flex-row items-center gap-5">
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

export default Header
