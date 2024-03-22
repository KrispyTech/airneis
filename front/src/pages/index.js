import Page from "@/components/Page"
import { useTranslation } from "next-i18next"
import { serverSideTranslations } from "next-i18next/serverSideTranslations"
import Link from "next/link"
import { useRouter } from "next/router"

export const getServerSideProps = async ({ locale }) => ({
  props: {
    ...(await serverSideTranslations(locale, ["common"]))
  }
})
const Home = () => {
  const { locale } = useRouter()
  const { t } = useTranslation("common")

  return (
    <Page className="h-screen text-secondary">
      <h2 className="text-2xl font-semibold py-5 px-3"> {t("homePage.wip")}</h2>
      <Link href="/" locale={locale}>
        {t("homePage.linkExample")}
      </Link>
    </Page>
  )
}

export default Home
