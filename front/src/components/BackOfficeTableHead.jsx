import { StopIcon } from "@heroicons/react/24/outline"

const BackOfficeTableHead = ({ handleChange, page }) => (
  <thead className="bg-clay">
    <tr>
      <th className="w-12 pl-2">
        <button onClick={handleChange} data-id="all" className="flex w-full">
          <StopIcon className="w-6 h-6" />
        </button>
      </th>
      <th className="text-start">IMAGES</th>
      <th className="text-start">NAME</th>
      {page === "categories" ? null : <th className="text-end w-20">PRICE</th>}
      <th></th>
    </tr>
  </thead>
)

export default BackOfficeTableHead
