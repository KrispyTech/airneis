import { StopIcon } from "@heroicons/react/24/outline"

const BackOfficeTableHead = ({ handleChange, columns }) => (
  <thead className="bg-clay">
    <tr>
      <th className="w-12 pl-2 py-1.5">
        <button onClick={handleChange} data-id="all" className="flex w-full">
          <StopIcon className="w-6 h-6" />
        </button>
      </th>
      {columns.map((column, i) => (
        <td
          key={`th-${i}`}
          className={`text-sm md:text-md lg:text-lg text-center ${column === "HTPrice" ? "text-end w-16 sm:w-20 md:w-24 lg:w-28" : ""} ${
            column === "name" ? "text-start" : ""
          }  uppercase font-bold `}
        >
          {column}
        </td>
      ))}
      <th></th>
    </tr>
  </thead>
)

export default BackOfficeTableHead
