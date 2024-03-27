import BackOfficeModal from "@/components/BackOfficeModal"
import BackOfficeTableHead from "@/components/BackOfficeTableHead"
import TableListItem from "@/components/TableListItem"
import routes from "@/routes"
import { PlusCircleIcon } from "@heroicons/react/24/outline"
import Link from "next/link"

const BackOfficeBoard = ({ handleOpenModal, itemId, isModalOpen, handleInput, handleChange, selected, items, columns, pageName }) => (
  <section className="flex flex-col w-full px-4 pt-3 overflow-y-scroll">
    {isModalOpen ? <BackOfficeModal handleOpenModal={handleOpenModal} id={itemId} /> : ""}
    <div className="flex justify-between">
      <h1 className="font-bold text-2xl">{pageName}</h1>
      <input placeholder="Search..." className="rounded-md drop-shadow-lg pl-2 w-40 md:w-72" onChange={handleInput} />
    </div>
    <Link href={routes.backoffice.addProduct} className="mt-4 py-0.5 px-1.5 bg-gray w-fit rounded-full border flex gap-1 items-center">
      <PlusCircleIcon className="w-6 h-6" />
      Add
    </Link>
    <table className="mt-4">
      <BackOfficeTableHead handleChange={handleChange} columns={columns} />
      <tbody>
        {items.map(item => (
          <TableListItem
            key={item.ID}
            item={item}
            selected={selected[item.ID]}
            handleOpenModal={handleOpenModal}
            handleChange={handleChange}
            columns={columns}
          />
        ))}
      </tbody>
    </table>
  </section>
)

export default BackOfficeBoard