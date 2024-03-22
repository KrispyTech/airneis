import BackOfficeListItem from "@/components/BackOfficeListItem"
import { ClockIcon, QuestionMarkCircleIcon, RectangleGroupIcon, ShoppingCartIcon, TagIcon, UserGroupIcon } from "@heroicons/react/24/outline"
import { useRouter } from "next/router"

const BackOfficeNav = () => {
  const router = useRouter()
  const [page] = router.pathname.split("/").slice(-1)

  return (
    <nav className={`h-full w-fit md:w-2/5 lg:w-1/3 xl:w-1/4 bg-tertiary border-r-2 z-20 ease-in-out duration-300`}>
      <ul className="flex flex-col px-3 border-b-2 gap-2 pb-3 ">
        <li className="md:text-xl border-b md:leading-extra-loose leading-extra-loose hidden md:block">Product Management</li>
        <li className="md:text-xl border-b-2 md:leading-extra-loose leading-extra-loose text-center font-bold md:hidden">PM</li>
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
      <ul className="flex flex-col px-3 gap-2 pb-6 ">
        <li className="md:text-xl border-b leading-extra-loose md:leading-extra-loose hidden md:block">User Management</li>
        <li className="md:text-xl border-b-2 md:leading-extra-loose leading-extra-loose text-center font-bold md:hidden">UM</li>
        <BackOfficeListItem name="users" page={page}>
          <UserGroupIcon className="h-8 w-8" strokeWidth="2.25" />
        </BackOfficeListItem>
        <BackOfficeListItem name="contacts" page={page}>
          <QuestionMarkCircleIcon className="h-8 w-8" strokeWidth="2.25" />
        </BackOfficeListItem>
      </ul>
    </nav>
  )
}

export default BackOfficeNav
