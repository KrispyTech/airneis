import BackOfficeListItem from "@/components/BackOfficeListItem"
import { ClockIcon, RectangleGroupIcon, ShoppingCartIcon, TagIcon } from "@heroicons/react/24/outline"

const BackOfficeNavPM = ({ page }) => (
  <ul className="flex flex-col px-3 border-b-2 gap-3 pb-3 ">
    <li className="md:text-xl border-b md:leading-extra-loose leading-extra-loose">Product Management</li>
    <BackOfficeListItem name="products" page={page}>
      <TagIcon className="h-8 w-8" strokeWidth="2.25" />
    </BackOfficeListItem>
    <BackOfficeListItem name="categories" page={page}>
      <RectangleGroupIcon className="h-8 w-8" strokeWidth="2.25" />
    </BackOfficeListItem>
    <BackOfficeListItem name="highlanders" page={page}>
      <ClockIcon className="h-8 w-8" strokeWidth="2.25" />
    </BackOfficeListItem>
    <BackOfficeListItem name="orders" page={page}>
      <ShoppingCartIcon className="h-8 w-8" strokeWidth="2.25" />
    </BackOfficeListItem>
  </ul>
)

export default BackOfficeNavPM
