import { EllipsisVerticalIcon, StopIcon } from "@heroicons/react/24/outline"
import { StopIcon as StopIconSolid } from "@heroicons/react/24/solid"

const TableListItem = ({ item, handleOpenModal, selected, handleChange, columns }) => (
  <tr className="odd:bg-listOdd even:bg-listEven">
    <td className="pl-2">
      <button className="flex" onClick={handleChange} data-id={item.ID}>
        {selected ? <StopIconSolid className="h-6 w-6" /> : <StopIcon className="w-6 h-6" />}
      </button>
    </td>
    {columns.map((column, i) => {
      let styles = ""
      let OTSText = ""

      if (column === "outOfStock") {
        if (item[column] === true) {
          styles = "text-red-400"
          OTSText = "out of stock"
        } else {
          styles = "text-green-400"
          OTSText = "in stock"
        }
      }

      return (
        <td
          key={`${item.ID}-${i}`}
          className={`text-center py-2 ${column === "name" ? "text-start" : ""} ${column === "priceWithoutTaxes" ? "text-end" : ""}`}
        >
          <span className={styles}>
            {column !== "outOfStock" ? item[column] : OTSText}
            {column === "priceWithoutTaxes" ? " €" : ""}
          </span>
        </td>
      )
    })}
    <td className="w-fit">
      <button className="flex justify-end w-full" onClick={handleOpenModal} data-id={item.id}>
        <EllipsisVerticalIcon className="w-6 h-6" />
      </button>
    </td>
  </tr>
)

export default TableListItem
