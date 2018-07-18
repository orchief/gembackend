select id,"from", "to", amount, nonce,fee,tx_hash,block_height,confirm_time,created,tx_state,is_token, 1 as o from tx as t1  where t1.from = "0xd6cb6744b7f2da784c5afd6b023d957188522198"  union all select id,"from", "to", amount, nonce,fee,tx_hash,block_height,confirm_time,created,tx_state,is_token, 0 as o from tx as t2 where t2.to = "0xd6cb6744b7f2da784c5afd6b023d957188522198" order by created desc ,id desc limit 10,10;

select id,tx_hash,`from`,`to`,amount,nonce,`decimal`,fee,block_height,confirm_time,created,log_index,tx_state,is_token, 0 as collection from token_tx as t1 where t1.from="0x0681d8db095565fe8a346fa0277bffde9c0edbbf" and t1.contract_addr=? union all select id,tx_hash,`from`,`to`,amount,nonce,`decimal`,fee,block_height,confirm_time,created,log_index,tx_state,is_token, 1 as collection from token_tx as t2 where t2.to="0x0681d8db095565fe8a346fa0277bffde9c0edbbf" and t2.contract_addr=? order by created desc,id desc limit 0,10;