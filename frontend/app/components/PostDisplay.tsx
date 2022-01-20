import { FC } from 'react'
import { Posts } from '~/routes'

interface Props {
  data: Posts
}

export const PostDisplay: FC<Props> = ({ data }) => {
  return (
    <div className="min-h-full flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div className="bg-white  shadow overflow-hidden sm:rounded-md">
          <ul role="list" className="divide-y divide-gray-200">
            {data.map((item) => (
              <li
                key={item.id}
                className="px-4 py-4 sm:px-6 hover:bg-gray-100 hover:shadow-lg"
              >
                {/* Your content */}
                <div className="flex">
                  <div>
                    <h4 className="text-lg font-bold">{item.title}</h4>
                    <p className="mt-1 prose">{item.description}</p>
                  </div>
                </div>
                <div>
                  {item.comments.map((comment) => (
                    <p key={comment.id} className="mt-1 prose">
                      {comment.text}
                    </p>
                  ))}
                </div>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  )
}
