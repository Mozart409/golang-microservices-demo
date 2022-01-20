import { ActionFunction, useLoaderData, Form, redirect } from 'remix'

import { Layout } from '~/components/Layout'
import { PostDisplay } from '~/components/PostDisplay'

import { LockClosedIcon } from '@heroicons/react/outline'
import { FC } from 'react'

export type Posts = Post[]

export interface Post {
  id: number
  title: string
  description: string
  comments: Comment[]
}

export interface Comment {
  id: number
  post_id: number
  text: string
}

export default function Index() {
  return (
    <div>
      <Layout>
        <h1 className="text-2xl font-semibold text-gray-900">
          Welcome to Remix
        </h1>
      </Layout>
    </div>
  )
}
