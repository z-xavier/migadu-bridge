import services from '@/services/alias';
import {
  ActionType,
  PageContainer,
  ProDescriptions,
  ProTable,
} from '@ant-design/pro-components';
import { Drawer } from 'antd';
import React, { useRef, useState } from 'react';
import { AliasItem } from '../../services/alias/typings';
import { useColumns } from './hooks/useColumns';

const { query } = services.AliasController;

const Alias: React.FC<unknown> = () => {
  const actionRef = useRef<ActionType>();
  const [row, setRow] = useState<AliasItem>();
  const { columns, checkColumns } = useColumns();

  return (
    <PageContainer
      header={{
        title: 'Token',
      }}
    >
      <ProTable<AliasItem>
        headerTitle="Alias"
        actionRef={actionRef}
        rowKey="id"
        search={{
          labelWidth: 'auto',
        }}
        editable={{
          type: 'multiple',
        }}
        request={async (params, sorter) => {
          let sortParams = {};
          const sorterKey = Object.keys(sorter ?? {})?.[0];
          if (sorterKey) {
            sortParams = {
              orderBy: `${sorterKey}:${
                sorter?.[sorterKey] === 'ascend' ? 'asc' : 'desc'
              }`,
            };
          }

          const res = await query({
            ...params,
            ...sortParams,
          });

          return {
            data: res?.data?.list || [],
            total: res?.data?.total,
          };
        }}
        columns={columns}
      />
      <Drawer
        width={600}
        open={!!row}
        onClose={() => {
          setRow(undefined);
        }}
        closable={false}
      >
        {row?.id && (
          <ProDescriptions<AliasItem>
            column={1}
            title={''}
            request={async () => ({
              data: row || {},
            })}
            params={{
              id: row?.id,
            }}
            columns={checkColumns}
          />
        )}
      </Drawer>
    </PageContainer>
  );
};

export default Alias;
