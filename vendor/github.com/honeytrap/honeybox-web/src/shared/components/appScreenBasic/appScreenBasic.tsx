import * as React from 'react';
import {
	Badge,
	Card,
	CardBody,
	CardHeader,
	Col,
	Pagination,
	PaginationItem,
	PaginationLink,
	Row,
	Table
} from 'reactstrap';
import { AppScreen } from '../appScreen/appScreen';

interface Props {
	title: string;
	children: any;
}

export function AppScreenBasic(props: Props) {
	return (
		<AppScreen>
			<Row>
				<Col>
					<Card>
						<CardHeader>{props.title}</CardHeader>
						<CardBody>
							{props.children}
						</CardBody>
					</Card>
				</Col>
			</Row>
		</AppScreen>
	)
}