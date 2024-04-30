import routes from "@/routes"
import clsx from "clsx"
import Link from "next/link"

const BackOfficeListItem = props => {
  const { children, name, page, ...otherProps } = props

  return (
    <li className={clsx("py-1.5 px-1.5", page === name ? "bg-black/30 rounded-sm " : "")} {...otherProps}>
      <Link href={routes.backoffice[name]} className="text-lg md:text-2xl flex items-center font-semibold gap-4 capitalize">
        {children}
        <span className="hidden md:block">{name}</span>
      </Link>
    </li>
  )
}

export default BackOfficeListItem
