import React from "react";
import { IDKitWidget, CredentialType } from "@worldcoin/idkit";

const colorButtonStyle = {
  padding: '10px 20px',
  border: 'none',
  cursor: 'pointer',
  backgroundColor: '#000',
  transition: '0.3s ease-in-out',
  color: 'white',
  borderRadius: '15px',
};

function WorldIdWidget() {

    const onSuccess = (result) => {
        console.log(result);
    };

    return ( 
		<div>
			<IDKitWidget 
				action = "loginVerify"
				signal = "my_signal"
				onSuccess = { onSuccess }
 				credential_types={[CredentialType.Orb,CredentialType.Phone]} 
				app_id = "app_staging_ef8611b8b7fc8451667562fbdae8c2e3"
				// walletConnectProjectId="get_this_from_walletconnect_portal"
			>	
			{({ open }) => <button onClick={open} style={colorButtonStyle}>Verify with WorldID</button>}	
			</IDKitWidget>
		 </div>
    );
}

export default WorldIdWidget;