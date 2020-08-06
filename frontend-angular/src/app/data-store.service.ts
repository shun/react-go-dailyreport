import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class DataStoreService {

  constructor() { }

  async postDate(data): Promise<boolean> {

    const param = {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    };

    return new Promise((resolve) => {
      fetch(
        `http://192.168.1.161:13000/users/${data.code}/reports/registry`,
        param
      )
      .then((res) => {
        console.log(res);
        if (!res.ok) {
          resolve(false);
        }
        resolve(true);
      })
    })
    .then((res: boolean) => {
      console.log(res);
      return res;
    });

  }
}
