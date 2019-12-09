import axios, { AxiosRequestConfig, AxiosInstance, AxiosError, AxiosResponse } from 'axios';
import { notification } from 'ant-design-vue';

const defaultAxiosConfig: AxiosRequestConfig = {
  baseURL: '/api',
};

export default class HttpClient {
  private axiosInstance: AxiosInstance;

  constructor(
    silent: boolean = true,
    axiosConfig: AxiosRequestConfig = defaultAxiosConfig,
  ) {
    this.axiosInstance = axios.create(axiosConfig);
    if (!silent) {
      this.axiosInstance.interceptors.response.use(
        (response) => response,
        (error) => {
          const { response, code } = error;
          if (code && code === 'ECONNABORTED') {
            notification.error({
              message: '请求超时',
              description: '请检查您的网络连接或与网站管理员联系',
            });
          } else if (response) {
            notification.error({
              message: '请求发生错误',
              description: response.data && response.data.errMsg,
            });
          } else {
            notification.error({
              message: '发生未知错误',
              description: '请联系网站管理员',
            });
          }
          return Promise.reject(response);
        },
      );
    }
  }

  public async get<T>(url: string, params?: any, headers?: any): Promise<T> {
    return (await this.axiosInstance.get<T>(url, { params, headers })).data;
  }

  public async delete<T>(url: string, params?: any, headers?: any): Promise<T> {
    return (await this.axiosInstance.delete<T>(url, { params, headers })).data;
  }

  public async head<T>(url: string, params?: any, headers?: any): Promise<T> {
    return (await this.axiosInstance.head<T>(url, { params, headers })).data;
  }

  public async post<T, B = any>(url: string, body: B, params?: any, headers?: any): Promise<T> {
    return (await this.axiosInstance.post<T>(url, body, { params, headers })).data;
  }

  public async put<T, B = any>(url: string, body: B, params?: any, headers?: any): Promise<T> {
    return (await this.axiosInstance.put<T>(url, body, { params, headers })).data;
  }

  public async patch<T, B = any>(url: string, body: B, params?: any, headers?: any): Promise<T> {
    return (await this.axiosInstance.patch<T>(url, body, { params, headers })).data;
  }
}

export const SILENT_HTTP_CLIENT = new HttpClient();
export const HTTP_CLIENT = new HttpClient(false);
