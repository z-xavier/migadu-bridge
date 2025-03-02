import services from '@/services/token';
import {
  ActionType,
  PageContainer,
  ProDescriptions,
  ProTable,
} from '@ant-design/pro-components';
import { Button, Drawer, message } from 'antd';
import dayjs from 'dayjs';
import React, { useRef, useState } from 'react';
import CustomModal from '../../components/CustomModal/CustomModal';
import {
  TokenAddRequest,
  TokenListItem,
  TokenUpdateRequest,
  TokenUpdateStatusRequest,
} from '../../services/token/typings';
import { useTokenColumns } from './hooks/useTokenColumns';

const { query, add, update, updateStatus, remove } = services.TokenController;

/**
 * 添加节点
 * @param fields
 */
const handleAdd = async (fields: TokenAddRequest) => {
  const hide = message.loading('正在添加');
  try {
    await add({ ...fields });
    hide();
    message.success('添加成功');
    return true;
  } catch (error) {
    hide();
    message.error('添加失败请重试！');
    return false;
  }
};

/**
 * 更新节点
 * @param fields
 */
const handleUpdate = async (fields: TokenUpdateRequest) => {
  const hide = message.loading('正在编辑');
  try {
    await update(fields.id, {
      description: fields.description,
      expiryAt: fields.expiryAt,
    });
    hide();
    message.success('编辑成功');
    return true;
  } catch (error) {
    hide();
    message.error('编辑失败请重试！');
    return false;
  }
};

const handleUpdateStatus = async (fields: TokenUpdateStatusRequest) => {
  const hide = message.loading('正在更新');
  try {
    await updateStatus(fields.id, fields.status);
    hide();
    message.success('更新成功');
    return true;
  } catch (error) {
    hide();
    message.error('更新失败请重试！');
    return false;
  }
};

/**
 *  删除节点
 * @param selectedRows
 */
const handleRemove = async (id: string) => {
  const hide = message.loading('正在删除');
  try {
    await remove(id);
    hide();
    message.success('删除成功，即将刷新');
    return true;
  } catch (error) {
    hide();
    message.error('删除失败，请重试');
    return false;
  }
};

const TokenList: React.FC<unknown> = () => {
  const [createModalVisible, setCreateModalVisible] = useState<boolean>(false);
  const [updateModalVisible, setUpdateModalVisible] = useState<boolean>(false);
  const [stepFormValues, setStepFormValues] = useState<TokenListItem>();
  const actionRef = useRef<ActionType>();
  const [row, setRow] = useState<TokenListItem>();
  const { columns, editColumns, checkColumns } = useTokenColumns({
    setUpdateModalVisible,
    setRow,
    setStepFormValues,
    handleUpdateStatus,
    handleRemove,
  });

  return (
    <PageContainer
      header={{
        title: 'Token',
      }}
    >
      <ProTable<TokenListItem>
        headerTitle="Token"
        actionRef={actionRef}
        rowKey="id"
        search={{
          labelWidth: 'auto',
        }}
        editable={{
          type: 'multiple',
        }}
        toolBarRender={() => [
          <Button
            key="1"
            type="primary"
            onClick={() => setCreateModalVisible(true)}
          >
            新建
          </Button>,
        ]}
        request={async (params, sorter, filter) => {
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
            // FIXME: remove @ts-ignore
            // @ts-ignore
            // filter,
          });

          return {
            data: res?.data?.list || [],
            total: res?.data?.total,
          };
        }}
        columns={columns}
      />
      <CustomModal
        onCancel={() => setCreateModalVisible(false)}
        modalVisible={createModalVisible}
      >
        <ProTable<TokenListItem, TokenAddRequest>
          onSubmit={async (value) => {
            let tempExpireAt;
            if (value?.expiryAt) {
              tempExpireAt = dayjs(value?.expiryAt).unix();
              const success = await handleAdd({
                ...value,
                expiryAt: tempExpireAt,
              });
              if (success) {
                setCreateModalVisible(false);
                if (actionRef.current) {
                  actionRef.current.reload();
                }
              }
            }
          }}
          rowKey="id"
          type="form"
          columns={columns}
        />
      </CustomModal>

      <CustomModal
        onCancel={() => setUpdateModalVisible(false)}
        modalVisible={updateModalVisible}
        title="编辑"
      >
        <ProTable<TokenListItem, Omit<TokenUpdateRequest, 'id'>>
          onSubmit={async (value) => {
            let tempExpireAt;
            if (value?.expiryAt && stepFormValues?.id) {
              tempExpireAt = dayjs(value?.expiryAt).unix();
              const success = await handleUpdate({
                id: stepFormValues?.id,
                description: value?.description,
                expiryAt: tempExpireAt,
              });
              if (success) {
                setUpdateModalVisible(false);
                if (actionRef.current) {
                  actionRef.current.reload();
                }
              }
            }
          }}
          rowKey="id"
          type="form"
          columns={editColumns}
          form={{
            initialValues: {
              ...stepFormValues,
              expiryAt: stepFormValues?.expiryAt
                ? dayjs.unix(stepFormValues?.expiryAt)
                : undefined,
            },
          }}
        />
      </CustomModal>

      <Drawer
        width={600}
        open={!!row}
        onClose={() => {
          setRow(undefined);
        }}
        closable={false}
      >
        {row?.id && (
          <ProDescriptions<TokenListItem>
            column={1}
            title=""
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

export default TokenList;
