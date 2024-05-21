import CategoryCard from "@/components/CategoryCard"
import Page from "@/components/Page"
import { getCategories } from "@/utils/services/categoryServices"
import { serverSideTranslations } from "next-i18next/serverSideTranslations"
import { useTranslation } from "next-i18next"

export const getServerSideProps = async ({ locale }) => {
  const { data: categories } = await getCategories()

  return {
    props: {
      categories,
      ...(await serverSideTranslations(locale, ["common"]))
    }
  }
}
const Home = props => {
  const { t } = useTranslation()
  const { categories } = props

  return (
    <Page className="h-screen text-secondary">
      <div className="w-full h-1/4 sm:h-1/3 bg-[url('https://imgs.search.brave.com/ZE_b8CzcWBoSZU5RW7btiybrw1zko5bsTjUxgo3MT1o/rs:fit:860:0:0/g:ce/aHR0cHM6Ly9mYWMu/aW1nLnBtZHN0YXRp/Yy5uZXQvc2NhbGUv/aHR0cC4zQS4yRi4y/RnByZDItYm9uZS1p/bWFnZS4yRXMzLXdl/YnNpdGUtZXUtd2Vz/dC0xLjJFYW1hem9u/YXdzLjJFY29tLjJG/RkFDLjJGMjAxOS4y/RjA5LjJGMjMuMkYw/YzQyMDU2Yi1lNjYx/LTQ2MmItODhiYi00/MWM3OTRlOTkyMjQu/MkVqcGVnL2F1dG94/NjAwL3F1YWxpdHkv/NjUvZ3JhbmQtY2Fu/YXBlLWRlLWphcmRp/bi5qcGc')] bg-center flex items-center justify-center">
        <h1 className="text-white drop-shadow-textOutline font-bold text-xl lg:text-3xl">{t("homePage.welcome")}</h1>
      </div>
      <p className="font-semibold text-3xl px-8 py-14 text-center">{t("homePage.description").toUpperCase()}</p>
      <ul className="grid grid-cols-3 gap-x-12 gap-y-8 px-8 pb-6">
        {categories.map(category => (
          <CategoryCard key={category.id} category={category} />
        ))}
      </ul>
    </Page>
  )
}

export default Home
