import { EllipsisVerticalIcon, StopIcon } from "@heroicons/react/24/outline"
import { StopIcon as StopIconSolid } from "@heroicons/react/24/solid"
import Image from "next/image"

const TableListItem = ({ product, handleOpenModal, selected, handleChange, columns }) => (
  <tr className="odd:bg-listOdd even:bg-listEven">
    <td className="pl-2">
      <button className="flex" onClick={handleChange} data-id={product.id}>
        {selected ? <StopIconSolid className="h-6 w-6" /> : <StopIcon className="w-6 h-6" />}
      </button>
    </td>
    {columns.map((column, i) => {
      let styles = ""

      if (column === "in stock") {
        if (product[column] === "Out of stock") {
          styles = "text-red-400"
        } else {
          styles = "text-green-400"
        }
      }

      return (
        <td
          key={`${product.id}-${i}`}
          className={`${column === "image" ? "py-1.5 hidden md:block w-24" : "text-center"} ${column === "name" ? "text-start" : ""} ${
            column === "HTPrice" ? "text-end" : ""
          }`}
        >
          {column === "image" ? (
            <Image src={product.image} alt={product.name} width={90} height={70} />
          ) : (
            <span className={styles}>
              {product[column]}
              {column === "HTPrice" ? " €" : ""}
            </span>
          )}
        </td>
      )
    })}
    <td className="w-fit">
      <button className="flex justify-end w-full" onClick={handleOpenModal} data-id={product.id}>
        <EllipsisVerticalIcon className="w-6 h-6" />
      </button>
    </td>
  </tr>
)

export default TableListItem
