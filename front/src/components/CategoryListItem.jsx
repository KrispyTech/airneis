import { EllipsisVerticalIcon, StopIcon } from "@heroicons/react/24/outline"
import { StopIcon as StopIconSolid } from "@heroicons/react/24/solid"
import Image from "next/image"

const CategoryListItem = ({ category, handleOpenModal, selected, handleChange }) => (
  <tr className="odd:bg-listOdd even:bg-listEven">
    <td className="pl-2">
      <button className="flex" onClick={handleChange} data-id={category.id}>
        {selected ? <StopIconSolid className="h-6 w-6" /> : <StopIcon className="w-6 h-6" />}
      </button>
    </td>
    <td>
      <Image src={category.image} alt="#" width={90} height={70} />
    </td>
    <td>{category.name}</td>
    <td className="w-fit">
      <button className="flex justify-end w-full" onClick={handleOpenModal} data-id={category.id}>
        <EllipsisVerticalIcon className="w-6 h-6" />
      </button>
    </td>
  </tr>
)

export default CategoryListItem
