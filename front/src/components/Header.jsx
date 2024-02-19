import SearchIcon from "@/components/icons/SearchIcon"
import ShoppingCartIcon from "@/components/icons/ShoppingCartIcon"
import BurgerMenuIcon from "@/components/icons/BurgerMenuIcon"

const Header = () => (
  <header className="bg-secondary text-primary px-4 flex justify-between items-center">
    <h1 className="text-4xl py-6">ÀIRNEIS</h1>
    <nav>
      <ul className="flex flex-row items-center gap-5">
        <li>
          <SearchIcon />
        </li>
        <li>
          <ShoppingCartIcon />
        </li>
        <li>
          <BurgerMenuIcon />
        </li>
      </ul>
    </nav>
  </header>
)

export default Header
