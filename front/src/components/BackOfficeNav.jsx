import BackOfficeNavPM from "@/components/BackOfficeNavPM"
import routes from "@/routes"
import { QuestionMarkCircleIcon, UserGroupIcon } from "@heroicons/react/24/outline"
import clsx from "clsx"
import Link from "next/link"
import { useRouter } from "next/router"

const BackOfficeNav = ({ isOpen, handleOpen }) => {
  const router = useRouter()
  const [page] = router.pathname.split("/").slice(-1)

  return (
    <nav
      className={`h-full w-fit md:left-0 md:w-2/5 lg:w-1/3 md:static bg-tertiary border-r-2 z-20 ${
        isOpen ? "left-0 ease-in-out duration-300 absolute" : "-left-48  ease-in-out duration-300 absolute"
      }`}
    >
      <BackOfficeNavPM page={page} />
      <ul className="flex flex-col px-3 gap-6 pb-6 ">
        <li className="md:text-xl border-b leading-extra-loose md:leading-extra-loose">User Management</li>
        <li className={clsx("py-1.5", page === "users" ? "bg-black/30 rounded-sm" : "")}>
          <Link href={routes.backoffice.users} className="text-lg md:text-2xl flex items-center font-bold gap-4">
            <UserGroupIcon className="h-8 w-8" strokeWidth="2.25" />
            Users
          </Link>
        </li>
        <li className={clsx("py-1.5", page === "contacts" ? "bg-black/30 rounded-sm" : "")}>
          <Link href={routes.backoffice.contacts} className="text-lg md:text-2xl flex items-center font-bold gap-4">
            <QuestionMarkCircleIcon className="h-8 w-8" strokeWidth="2.25" />
            Contacts
          </Link>
        </li>
        <button className="text-xl md:hidden mt-4 bg-black/10 border rounded-sm" onClick={handleOpen}>
          Close
        </button>
      </ul>
    </nav>
  )
}

export default BackOfficeNav
