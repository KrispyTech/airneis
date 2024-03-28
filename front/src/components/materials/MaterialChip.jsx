const MaterialChip = props => {
  const { name } = props

  return <div className="bg-gray rounded-full px-3 py-1 text-sm font-semibold border w-24 text-center">{name}</div>
}

export default MaterialChip
