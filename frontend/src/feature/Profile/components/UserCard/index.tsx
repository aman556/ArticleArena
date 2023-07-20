import React from 'react';
import userDefault from '../../../../constants/userDefault.png';
import useStyles from './styles';
import InfoPair from '../../../../components/InfoPair';
import { Button } from '../../../../styles/Button';

export interface IUserCard {
  name: string;
  designation: string;
  address: string;
}

const UserCard: React.FC<IUserCard> = (props) => {
  const styles = useStyles();
  const imageURL = 'https://images.unsplash.com/photo-1533035350251-aa8b8e208d95';

  return (
    <div className={styles.usercard}>
      <img src={userDefault} height={150} width={150} className={styles.imageStyles} alt='Kartikay ki fotu'/>
      <div className={styles.additionalText}>
        <InfoPair title={props.name} classes={{ title: styles.titleText }} />
        <InfoPair title={props.designation} />
        <InfoPair title={props.address} />
      </div>
      <div className={styles.cls_button}>
        <Button>Follow</Button>
        <Button >Message</Button>
      </div>
    </div>
  );
};

export default UserCard;
