import routes from "@/routes"
import { ClockIcon, RectangleGroupIcon, ShoppingCartIcon, TagIcon } from "@heroicons/react/24/outline"
import clsx from "clsx"
import Link from "next/link"

const BackOfficeNavPM = ({ page }) => (
  <ul className="flex flex-col px-3 border-b-2 gap-3 pb-3 ">
    <li className="md:text-xl border-b md:leading-extra-loose leading-extra-loose">Product Management</li>
    <li className={clsx("py-1.5", page === "products" ? "bg-black/30 rounded-sm" : "")}>
      <Link href={routes.backoffice.products} className="text-lg md:text-2xl flex items-center font-bold gap-4">
        <TagIcon className="h-8 w-8" strokeWidth="2.25" />
        Products
      </Link>
    </li>
    <li className={clsx("py-1.5", page === "categories" ? "bg-black/30 rounded-sm" : "")}>
      <Link href={routes.backoffice.categories} className="text-lg md:text-2xl flex items-center font-bold gap-4">
        <RectangleGroupIcon className="h-8 w-8" strokeWidth="2.25" />
        Categories
      </Link>
    </li>
    <li className={clsx("py-1.5", page === "highlanders" ? "bg-black/30 rounded-sm" : "")}>
      <Link href={routes.backoffice.highlanders} className="text-lg md:text-2xl flex items-center font-bold gap-4">
        <ClockIcon className="h-8 w-8" strokeWidth="2.25" />
        Highlanders
      </Link>
    </li>
    <li className={clsx("py-1.5", page === "orders" ? "bg-black/30 rounded-sm" : "")}>
      <Link href={routes.backoffice.orders} className="text-lg md:text-2xl flex items-center font-bold gap-4">
        <ShoppingCartIcon className="h-8 w-8" strokeWidth="2.25" />
        Orders
      </Link>
    </li>
  </ul>
)

export default BackOfficeNavPM
