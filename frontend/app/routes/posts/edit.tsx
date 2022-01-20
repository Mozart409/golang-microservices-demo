import { Layout } from '~/components/Layout'

export async function createComment(formData: FormData) {
  const comment = await fetch('http://localhost:4000/api/comment', {
    method: 'POST',
    body: formData,
  })
  return comment
}

export default function Edit() {
  return (
    <div>
      <Layout>
        <h1 className="text-2xl font-semibold text-gray-900">Edit</h1>
      </Layout>
    </div>
  )
}
