import React, {useEffect, useState} from 'react'
import {Card, Space} from 'antd'
import {ListHotTopic} from "../service";
import {useHistory} from "react-router-dom";
import styles from './hotTopic.module.css'
import {configResponsive, useResponsive} from 'ahooks';

configResponsive({
    small: 0,
    middle: 800,
    large: 1200,
});
const HostTopic: React.FC = () => {
    const [data, setData] = useState([])
    const history = useHistory()

    const responsive = useResponsive();

    useEffect(() => {
        ListHotTopic().then(res => {
            console.log(res)
            if (res.ok) {
                setData(res.data)
            }
        })
    }, [])
    const toHot = function (item) {
        console.log(item)
        let id = item.shortCode == "" ? item.id : item.shortCode
        history.push(`/t/${id}`)
    }

    return (
        <>
            <Card title="发现" className={styles.title}>

                <Space className={styles.container} wrap={true}
                       direction={responsive['middle'] ? "horizontal" : "vertical"}>
                    {
                        data.map((item) =>
                            <Card hoverable={true} key={"card" + item.id} className={styles.card} onClick={() => {
                                toHot(item);
                            }}>
                                <span>{item.name}</span>
                            </Card>
                        )
                    }
                </Space>


            </Card>
        </>

    )
}


export default HostTopic
