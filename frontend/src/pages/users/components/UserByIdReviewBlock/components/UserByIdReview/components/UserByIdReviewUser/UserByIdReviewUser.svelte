<script lang="ts">
  import type { User } from '@/types/User.ts'
  import UserByIdReviewUserLayout from '@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewUser/UserByIdReviewUserLayout.svelte'
  import { Link } from 'svelte-routing'
  import Avatar from '@/common-components/Avatar/Avatar.svelte'

  type UserByIdReviewProps = User & {
    user_id?: number
  }

  const { avatar, username, id, user_id }: UserByIdReviewProps = $props()
</script>

<UserByIdReviewUserLayout>
  <Link to={`/users/${id === undefined ? user_id : id}`}>
    <Avatar src={avatar} alt="img" size={'extra small'} />
  </Link>
  <Link to={`/users/${id === undefined ? user_id : id}`}>
    <p class="hover:underline ease-out duration-200">{username}</p>
  </Link>
</UserByIdReviewUserLayout>
