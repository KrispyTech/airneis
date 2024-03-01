import { MagnifyingGlassIcon, ShoppingCartIcon, Bars3Icon } from "@heroicons/react/24/outline"
import webConfig from "@@/webConfig"

const Header = () => (
  <header className="bg-secondary text-primary px-4 flex justify-between items-center">
    <h1 className="text-4xl py-6">ÀIRNEIS</h1>
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
  </header>
)

export default Header
