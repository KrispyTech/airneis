import routes from "@/routes"
import { CheckCircleIcon, PencilIcon, TrashIcon, XCircleIcon } from "@heroicons/react/24/outline"
import Link from "next/link"
import { useState } from "react"

// eslint-disable-next-line no-unused-vars
const BackOfficeModal = ({ handleOpenModal, id }) => {
  const [askDelete, setAskDelete] = useState(false)
  const handleAskDelete = () => setAskDelete(!askDelete)
  const handleDelete = () => {
    // To do
  }

  return (
    <div className="absolute inset-0 flex flex-col items-center justify-center bg-black/20 z-10">
      <div className="bg-modal w-2/3 md:w-1/3 lg:w-1/4 h-1/4 flex flex-col items-center justify-evenly">
        {askDelete ? (
          <>
            <h1 className="border-b w-5/6 text-center text-2xl md:text-4xl">DELETE ?</h1>
            <button className="flex text-green-800 bg-gray border rounded-full w-fit px-2 py-1 gap-1" onClick={handleDelete}>
              <CheckCircleIcon className="h-6 w-6 text-green-400" />
              Yes
            </button>
            <button className="flex text-red-800 bg-gray border rounded-full w-fit px-2 py-1 gap-1" onClick={handleAskDelete}>
              <XCircleIcon className="h-6 w-6 text-red-400" />
              No
            </button>
            <button className="hover:underline" onClick={handleOpenModal}>
              Back
            </button>
          </>
        ) : (
          <>
            <h1 className="border-b w-5/6 text-center text-2xl md:text-4xl">ACTIONS</h1>
            <Link href={routes.home} className="flex text-green-800 bg-gray border rounded-full w-fit px-2 py-1 gap-1">
              <PencilIcon className="h-6 w-6 text-green-400" />
              Update
            </Link>
            <button className="flex text-red-800 bg-gray border rounded-full w-fit px-2 py-1 gap-1" onClick={handleAskDelete}>
              <TrashIcon className="h-6 w-6 text-red-400" />
              Delete
            </button>
            <button className="hover:underline" onClick={handleOpenModal}>
              Back
            </button>
          </>
        )}
      </div>
    </div>
  )
}

export default BackOfficeModal
