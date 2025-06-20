<script>
  import AppLayout from '@/AppLayout.svelte'
  import NotFound from '@/common-components/NotFound/NotFound.svelte'
  import ConfirmEmailPage from '@/pages/confirm-email/ConfirmEmailPage.svelte'
  import ForgotPasswordPage from '@/pages/forgot-password/ForgotPasswordPage.svelte'
  import OrderConfirmPayPage from '@/pages/order-confirm-pay/OrderConfirmPayPage.svelte'
  import OrderRequestPage from '@/pages/order-request/OrderRequestPage.svelte'
  import OrderSubmitRequirementsPage from '@/pages/order-submit-requirements/OrderSubmitRequirementsPage.svelte'
  import ResetPasswordPage from '@/pages/reset-password/ResetPasswordPage.svelte'
  import SearchPage from '@/pages/search/SearchPage.svelte'
  import ServicePage from '@/pages/services/ServicePage.svelte'
  import SignInPage from '@/pages/sign-in/SignInPage.svelte'
  import SignUpPage from '@/pages/sign-up/SignUpPage.svelte'
  import { onDestroy, onMount } from 'svelte'
  import { Route, Router } from 'svelte-routing'
  import { connectToWebsocket } from './api/websocket'
  import { signDataStore } from './common-components/Header/components/HeaderProfileBlock/sign-data-store'
  import Home from './pages/home/HomePage.svelte'
  import MyProfileEditPage from './pages/my-profile/sub-pages/edit/MyProfileEditPage.svelte'
  import MyProfileOrdersPage from './pages/my-profile/sub-pages/orders/MyProfileOrdersPage.svelte'
  import MyProfileOrderByIdPage from './pages/my-profile/sub-pages/orders/sub-pages/order-by-id/MyProfileOrderByIdPage.svelte'
  import MyProfileRequestsPage from './pages/my-profile/sub-pages/requests/MyProfileRequestsPage.svelte'
  import MyProfileServicesPage from './pages/my-profile/sub-pages/services/MyProfileServicesPage.svelte'
  import CreateServiceBasicInfoPage from './pages/my-profile/sub-pages/services/sub-pages/create-service-basic-info/CreateServiceBasicInfoPage.svelte'
  import CreateServiceMediaPage from './pages/my-profile/sub-pages/services/sub-pages/create-service-media/CreateServiceMediaPage.svelte'
  import CreateServicePackagesPage from './pages/my-profile/sub-pages/services/sub-pages/create-service-packages/CreateServicePackagesPage.svelte'
  import CreateServiceRequirementsPage from './pages/my-profile/sub-pages/services/sub-pages/create-service-requirements/CreateServiceRequirementsPage.svelte'
  import UserPage from './pages/users/UserPage.svelte'
  import { openIndexedDb } from './utils/indexed-db-utils'
  import { getUserSession } from './common-components/Header/helpers'
  import MyProfileRequestByIdPage from './pages/my-profile/sub-pages/requests/sub-pages/request-by-id/MyProfileRequestByIdPage.svelte'

  let mounted = $state(false)

  const unsubscribe = signDataStore.subscribe((value) => {
    if (value && value.authenticated) connectToWebsocket()
  })

  onMount(() => {
    openIndexedDb()
    getUserSession()
      .then((resp) => {
        signDataStore.set(resp)
      })
      .finally(() => {
        mounted = true
      })
  })

  onDestroy(() => unsubscribe())
</script>

{#if mounted}
  <Router>
    <AppLayout>
      <Route path="/">
        <Home />
      </Route>
      <Route path="/services/:id" let:params>
        <ServicePage id={params.id} />
      </Route>
      <Route path="/orders/request">
        <OrderRequestPage />
      </Route>
      <Route path="/orders/confirm-pay">
        <OrderConfirmPayPage />
      </Route>
      <Route path="/orders/:orderId/submit-requirements" let:params>
        <OrderSubmitRequirementsPage orderId={params.orderId} />
      </Route>
      <Route path="/users/:id" let:params>
        <UserPage id={params.id} />
      </Route>
      <Route path="/search">
        <SearchPage />
      </Route>
      <Route path="/sign-in">
        <SignInPage />
      </Route>
      <Route path="/sign-up">
        <SignUpPage />
      </Route>
      <Route path="/confirm-email">
        <ConfirmEmailPage />
      </Route>
      <Route path="/forgot-password">
        <ForgotPasswordPage />
      </Route>
      <Route path="/reset-password">
        <ResetPasswordPage />
      </Route>
      <Route path="/my-profile/orders">
        <MyProfileOrdersPage />
      </Route>
      <Route path="/my-profile/services">
        <MyProfileServicesPage />
      </Route>
      <Route path="/my-profile/requests">
        <MyProfileRequestsPage />
      </Route>
      <Route path="/my-profile/orders/:orderId" let:params>
        <MyProfileOrderByIdPage orderId={params.orderId} />
      </Route>
      <Route path="/my-profile/requests/:orderId" let:params>
        <MyProfileRequestByIdPage orderId={params.orderId} />
      </Route>
      <Route path="/my-profile/services/create/basic-information">
        <CreateServiceBasicInfoPage />
      </Route>
      <Route path="/my-profile/services/create/packages">
        <CreateServicePackagesPage />
      </Route>
      <Route path="/my-profile/services/create/requirements">
        <CreateServiceRequirementsPage />
      </Route>
      <Route path="/my-profile/services/create/media">
        <CreateServiceMediaPage />
      </Route>
      <Route path="/my-profile/edit">
        <MyProfileEditPage />
      </Route>
      <Route path="/*">
        <NotFound />
      </Route>
    </AppLayout>
  </Router>
{/if}
