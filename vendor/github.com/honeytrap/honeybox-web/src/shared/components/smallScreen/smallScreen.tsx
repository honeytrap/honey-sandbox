import * as React from 'react';
import { Button, Card, CardBody, CardGroup, Col, Container, Input, InputGroup, InputGroupAddon, InputGroupText, Row } from 'reactstrap';

interface Props {
	children: any;
}

export function SmallScreen(props: Props) {
	return (
		<div className="app flex-row align-items-center">
			<Container>
				<Row className="justify-content-center">
					<Col md="6">
						<CardGroup>
							<Card className="mx-4">
								<CardBody>
									{props.children}
								</CardBody>
							</Card>
						</CardGroup>
					</Col>
				</Row>
			</Container>
		</div>
	)
}