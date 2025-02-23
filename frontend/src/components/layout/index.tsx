import { DashboardLayout } from '@toolpad/core'
import { NextAppProvider } from '@toolpad/core/nextjs'
import * as React from 'react'
import { NAVIGATION } from '../../constants/navigation'

interface LayoutProps {
  children?: React.ReactNode
}

export default function Layout({ children }: Readonly<LayoutProps>) {
  return (
    <NextAppProvider navigation={NAVIGATION}>
      <DashboardLayout
        slots={{
          appTitle: () => <div />,
        }}
      >
        {children}
      </DashboardLayout>
    </NextAppProvider>
  )
}
