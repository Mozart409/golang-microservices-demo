import { Layout } from '~/components/Layout'
import { PostForm } from '~/components/PostForm'

export default function Index() {
  return (
    <div>
      <Layout>
        <h1>Welcome to Remix</h1>
        <PostForm action={''} />
      </Layout>
    </div>
  )
}
